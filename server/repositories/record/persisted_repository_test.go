package record

import (
	"assets-liabilities/entities"
	"assets-liabilities/logging"
	"assets-liabilities/types"
	"context"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.Nil(t, err)
	db.AutoMigrate(&entities.User{}, &entities.Record{})
	return db
}

func initCtx(t *testing.T) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, logging.Ctx{}, logrus.New())
	return ctx
}

func TestPersistedRepository_FindByID(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  uint64
	}

	db := initDB(t)
	defer db.Close()

	ctx := initCtx(t)

	r := &PersistedRepository{
		db: db,
	}
	asset1, err := r.Create(ctx, entities.Record{
		Name:    "Asset1",
		Type:    entities.Asset,
		Balance: 120.45,
	})

	require.Nil(t, err)
	liability1, err := r.Create(ctx, entities.Record{
		Name:    "Liability1",
		Type:    entities.Liability,
		Balance: 220.45,
	})
	require.Nil(t, err)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Record
		wantErr bool
	}{
		{
			name: "Unknown ID",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  23132131,
			},
			want:    entities.Record{},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  12,
			},
			want:    entities.Record{},
			wantErr: true,
		},
		{
			name: "Find Asset",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  asset1.ID,
			},
			want:    asset1,
			wantErr: false,
		},
		{
			name: "Find Liability",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  liability1.ID,
			},
			want:    liability1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PersistedRepository{
				db: tt.fields.db,
			}
			got, err := r.FindByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PersistedRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PersistedRepository.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersistedRepository_List(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		where  *entities.Record
		params *entities.QueryParams
	}

	ctx := initCtx(t)

	db := initDB(t)
	defer db.Close()

	r := &PersistedRepository{
		db: db,
	}
	asset1, err := r.Create(ctx, entities.Record{
		Name:    "Asset1",
		Type:    entities.Asset,
		Balance: 120.45,
	})
	require.Nil(t, err)

	asset2, err := r.Create(ctx, entities.Record{
		Name:    "Asset2",
		Type:    entities.Asset,
		Balance: 0.12,
	})
	require.Nil(t, err)

	liability1, err := r.Create(ctx, entities.Record{
		Name:    "Liability1",
		Type:    entities.Liability,
		Balance: 220.45,
	})
	require.Nil(t, err)

	liability2, err := r.Create(ctx, entities.Record{
		Name:    "Liability2",
		Type:    entities.Liability,
		Balance: 92.45,
	})
	require.Nil(t, err)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Record
		wantErr bool
	}{
		{
			name: "List All",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
			},
			want: []entities.Record{
				asset1,
				asset2,
				liability1,
				liability2,
			},
			wantErr: false,
		},
		{
			name: "List All Assets",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				where: &entities.Record{
					Type: entities.Asset,
				},
			},
			want: []entities.Record{
				asset1,
				asset2,
			},
			wantErr: false,
		},
		{
			name: "List All Liabilities",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				where: &entities.Record{
					Type: entities.Liability,
				},
			},
			want: []entities.Record{
				liability1,
				liability2,
			},
			wantErr: false,
		},
		{
			name: "List All Assets With Limit 1",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				where: &entities.Record{
					Type: entities.Asset,
				},
				params: &entities.QueryParams{
					Limit: types.CreateInt(1),
				},
			},
			want: []entities.Record{
				asset1,
			},
			wantErr: false,
		},
		{
			name: "List All Assets With Limit Higher than Total Assets",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				where: &entities.Record{
					Type: entities.Asset,
				},
				params: &entities.QueryParams{
					Limit: types.CreateInt(3),
				},
			},
			want: []entities.Record{
				asset1,
				asset2,
			},
			wantErr: false,
		},
		{
			name: "List All Assets With Limit and Offset",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				where: &entities.Record{
					Type: entities.Asset,
				},
				params: &entities.QueryParams{
					Limit:  types.CreateInt(1),
					Offset: types.CreateInt(1),
				},
			},
			want: []entities.Record{
				asset2,
			},
			wantErr: false,
		},
		{
			name: "List All with Out of Range Limit and Offset",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				where: &entities.Record{
					Type: entities.Asset,
				},
				params: &entities.QueryParams{
					Limit:  types.CreateInt(100),
					Offset: types.CreateInt(100),
				},
			},
			want:    []entities.Record{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PersistedRepository{
				db: tt.fields.db,
			}
			got, err := r.List(tt.args.ctx, tt.args.where, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("PersistedRepository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PersistedRepository.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersistedRepository_Create(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data entities.Record
	}

	ctx := initCtx(t)

	db := initDB(t)
	defer db.Close()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Record
		wantErr bool
	}{
		{
			name: "Create Asset",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				data: entities.Record{
					Name:    "NewAsset",
					Type:    entities.Asset,
					Balance: 500.01,
				},
			},
			want: entities.Record{
				Name:    "NewAsset",
				Type:    entities.Asset,
				Balance: 500.01,
			},
			wantErr: false,
		},
		{
			name: "Create Liability",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				data: entities.Record{
					Name:    "NewLiability",
					Type:    entities.Liability,
					Balance: 100.033,
				},
			},
			want: entities.Record{
				Name:    "NewLiability",
				Type:    entities.Liability,
				Balance: 100.033,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PersistedRepository{
				db: tt.fields.db,
			}
			got, err := r.Create(tt.args.ctx, tt.args.data)

			got.ID = tt.want.ID
			got.CreatedAt = tt.want.CreatedAt
			got.UpdatedAt = tt.want.UpdatedAt

			if (err != nil) != tt.wantErr {
				t.Errorf("PersistedRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PersistedRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersistedRepository_Update(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data entities.Record
	}

	ctx := initCtx(t)

	db := initDB(t)
	defer db.Close()

	r := &PersistedRepository{
		db: db,
	}

	asset1, err := r.Create(ctx, entities.Record{
		Name:    "Asset1",
		Type:    entities.Asset,
		Balance: 120.45,
	})
	require.Nil(t, err)

	asset2, err := r.Create(ctx, entities.Record{
		Name:    "Asset2",
		Type:    entities.Asset,
		Balance: 0.12,
	})
	require.Nil(t, err)

	_, err = r.Create(ctx, entities.Record{
		Name:    "Liability1",
		Type:    entities.Liability,
		Balance: 220.45,
	})
	require.Nil(t, err)

	_, err = r.Create(ctx, entities.Record{
		Name:    "Liability2",
		Type:    entities.Liability,
		Balance: 92.45,
	})
	require.Nil(t, err)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Record
		wantErr bool
	}{
		{
			name: "Update Asset",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				data: entities.Record{
					Entity: entities.Entity{
						ID: asset2.ID,
					},
					Name:    "UpdatedAsset2",
					Balance: 150.23,
				},
			},
			want: entities.Record{
				Entity: entities.Entity{
					ID:        asset2.ID,
					CreatedAt: asset2.CreatedAt,
				},
				Name:    "UpdatedAsset2",
				Type:    entities.Asset,
				Balance: 150.23,
			},
			wantErr: false,
		},
		{
			name: "Change Record Type",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				data: entities.Record{
					Entity: entities.Entity{
						ID: asset1.ID,
					},
					Type: entities.Liability,
				},
			},
			want: entities.Record{
				Entity: entities.Entity{
					ID:        asset1.ID,
					CreatedAt: asset1.CreatedAt,
				},
				Name:    "Asset1",
				Type:    entities.Liability,
				Balance: 120.45,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PersistedRepository{
				db: tt.fields.db,
			}
			got, err := r.Update(tt.args.ctx, tt.args.data)
			tt.want.UpdatedAt = got.UpdatedAt
			if (err != nil) != tt.wantErr {
				t.Errorf("PersistedRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PersistedRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersistedRepository_Delete(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  uint64
	}

	ctx := initCtx(t)

	db := initDB(t)
	defer db.Close()

	r := &PersistedRepository{
		db: db,
	}

	asset1, err := r.Create(ctx, entities.Record{
		Name:    "Asset1",
		Type:    entities.Asset,
		Balance: 120.45,
	})
	require.Nil(t, err)

	asset2, err := r.Create(ctx, entities.Record{
		Name:    "Asset2",
		Type:    entities.Asset,
		Balance: 0.12,
	})
	require.Nil(t, err)

	_, err = r.Create(ctx, entities.Record{
		Name:    "Liability1",
		Type:    entities.Liability,
		Balance: 220.45,
	})
	require.Nil(t, err)

	_, err = r.Create(ctx, entities.Record{
		Name:    "Liability2",
		Type:    entities.Liability,
		Balance: 92.45,
	})
	require.Nil(t, err)

	beforeDelete, err := r.List(ctx, nil, nil)
	require.Nil(t, err)
	assert.Len(t, beforeDelete, 4)

	assert.Nil(t, r.Delete(ctx, asset2.ID))

	afterDelete, err := r.List(ctx, nil, nil)
	require.Nil(t, err)

	assert.Len(t, afterDelete, 3)

	updatedAssets, err := r.List(ctx, &entities.Record{
		Type: entities.Asset,
	}, nil)
	require.Nil(t, err)
	require.Len(t, updatedAssets, 1)
	assert.Equal(t, asset1.ID, updatedAssets[0].ID)

}
