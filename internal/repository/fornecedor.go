package repository

import (
	"strconv"
	"time"

	"github.com/ericNKS/estoque/internal/db"
	"github.com/ericNKS/estoque/internal/entities/types"
	"gorm.io/gorm"
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
	err := fr.rdb.Set(strconv.FormatUint(uint64(f.ID), 10), f, 20*time.Hour)
	if err != nil {
		return err
	}
	return nil
}

func (fr *FornecedorRepository) FindAll() ([]*types.Fornecedor, error) {
	db, err := db.Connection()
	if err != nil {
		return nil, err
	}

	var fornecedores []*types.Fornecedor
	result := db.Find(&fornecedores)
	if result.Error != nil {
		return nil, err
	}

	return fornecedores, nil
}

func (fr *FornecedorRepository) IsUnique(cnpj string, instituicaoId uint64) (bool, error) {
	db, err := db.Connection()
	if err != nil {
		return false, err
	}
	f := &types.Fornecedor{}
	result := db.First(f, "cnpj = ? AND instituicao_id = ?", cnpj, instituicaoId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return true, nil
		}
		return false, result.Error
	}
	return false, nil
}

func (fr *FornecedorRepository) FindById(id uint64) (*types.Fornecedor, error) {
	defer fr.rdb.Db.Close()
	var f types.Fornecedor
	err := fr.rdb.Get(strconv.FormatUint(id, 10), &f)
	if err == nil {
		return &f, nil
	}

	result := fr.db.First(&f, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	fr.rdb.Set(strconv.FormatUint(uint64(f.ID), 10), &f, 20*time.Hour)

	return &f, nil
}

func (fr *FornecedorRepository) Update(f *types.Fornecedor) error {
	db, err := db.Connection()
	if err != nil {
		return err
	}

	db.Save(f)
	return nil
}

func (fr *FornecedorRepository) Delete(f *types.Fornecedor) error {
	db, err := db.Connection()
	if err != nil {
		return err
	}

	db.Delete(f)
	return nil
}
