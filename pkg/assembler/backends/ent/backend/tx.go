//
// Copyright 2023 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backend

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/guacsec/guac/pkg/assembler/backends/ent"
)

func WithinTX[T any](ctx context.Context, entClient *ent.Client, exec func(ctx context.Context) (*T, error)) (*T, error) {
	if entClient == nil {
		return nil, Errorf("%v ::  %s", "WithinTX", "ent client is not initialized")
	}

	tx, err := entClient.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		fmt.Printf("BeginTx err %s for tx %v ctx %v\n", err, &tx, &ctx)
		return nil, err
	}
	fmt.Printf("BeginTx tx %v %v\n", &tx, &ctx)

	ctx = ent.NewTxContext(ctx, tx)

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		}
	}()

	result, err := exec(ctx)
	if err != nil {
		fmt.Printf("exec err %s for tx %v ctx %v\n", err, &tx, &ctx)
		errRollback := tx.Rollback()
		fmt.Printf("exec errRollback %s for tx %v ctx %v\n", errRollback, &tx, &ctx)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		fmt.Printf("Commit err %s %v ctx %v\n", err, &tx, &ctx)
		return nil, err
	}

	fmt.Printf("tx commit %v ctx %v\n", &tx, &ctx)
	return result, nil
}
