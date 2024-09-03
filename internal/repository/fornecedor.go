package repository

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ericNKS/estoque/internal/db"
	"github.com/ericNKS/estoque/internal/entities/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FornecedorRepository struct {
	db  *gorm.DB
	rdb *db.DbRedis
}

func NewFornecedorRepository() (*FornecedorRepository, error) {
	pstg, err := db.Connection()
	if err != nil {
		return nil, err
	}
	rdb := db.RedisConnection()

	return &FornecedorRepository{
		db:  pstg,
		rdb: rdb,
	}, nil
}

func (fr *FornecedorRepository) Create(f *types.Fornecedor) error {
	result := fr.db.Create(f)
	if result.Error != nil {
		return result.Error
	}
	go fr.rdb.Set(strconv.FormatUint(uint64(f.ID), 10), f, 20*time.Hour)
	defer fr.rdb.Db.Close()

	return nil
}

func (fr *FornecedorRepository) FindAll() ([]*types.Fornecedor, error) {
	var fornecedores []*types.Fornecedor
	result := fr.db.
		Limit(10).
		Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: "nome_fantasia",
			}, Desc: false,
		}).
		Find(&fornecedores)

	if result.Error != nil {
		return nil, result.Error
	}

	return fornecedores, nil
}

func (fr *FornecedorRepository) IsUnique(cnpj string, instituicaoId uint64) error {
	f := &types.Fornecedor{}
	result := fr.db.First(f, "cnpj = ? AND instituicao_id = ?", cnpj, instituicaoId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil
		}
		return result.Error
	}
	return nil
}

func (fr *FornecedorRepository) FindById(id uint64) (*types.Fornecedor, error) {
	var f types.Fornecedor
	err := fr.rdb.Get(strconv.FormatUint(id, 10), &f)
	if err == nil {
		fmt.Println("Redis")
		return &f, nil
	}

	result := fr.db.First(&f, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	fr.rdb.Set(strconv.FormatUint(uint64(f.ID), 10), &f, 20*time.Hour)
	defer fr.rdb.Db.Close()

	return &f, nil
}

func (fr *FornecedorRepository) Update(f *types.Fornecedor) error {
	fr.db.Save(f)
	go fr.rdb.Set(strconv.FormatUint(uint64(f.ID), 10), &f, 20*time.Hour)
	defer fr.rdb.Db.Close()

	return nil
}

func (fr *FornecedorRepository) Delete(f *types.Fornecedor) error {
	defer fr.rdb.Db.Close()
	err := fr.rdb.Del(strconv.FormatUint(uint64(f.ID), 10))
	if err != nil {
		return err
	}

	result := fr.db.Delete(f)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
