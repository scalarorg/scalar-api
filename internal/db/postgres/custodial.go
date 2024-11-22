package postgres

import "github.com/scalarorg/xchains-api/internal/db/postgres/models"

func (s *ScalarClient) SaveCustodial(custodial *models.Custodial) error {
	return s.scalarPostgresClient.Db.Create(custodial).Error
}

func (s *ScalarClient) GetCustodials() ([]*models.Custodial, error) {
	var custodials []*models.Custodial
	if err := s.scalarPostgresClient.Db.Find(&custodials).Error; err != nil {
		return nil, err
	}
	return custodials, nil
}

func (s *ScalarClient) GetCustodialByName(name string) (*models.Custodial, error) {
	var custodial models.Custodial
	if err := s.scalarPostgresClient.Db.Where("name = ?", name).First(&custodial).Error; err != nil {
		return nil, err
	}
	return &custodial, nil
}

func (s *ScalarClient) GetCustodialByNames(names []string) ([]*models.Custodial, error) {
	var custodials []*models.Custodial
	if err := s.scalarPostgresClient.Db.Where("name IN (?)", names).Find(&custodials).Error; err != nil {
		return nil, err
	}
	return custodials, nil
}

func (s *ScalarClient) SaveCustodialGroup(custodialGroup *models.CustodialGroup) error {
	return s.scalarPostgresClient.Db.Create(custodialGroup).Error
}

func (s *ScalarClient) GetCustodialGroups() ([]*models.CustodialGroup, error) {
	var custodialGroups []*models.CustodialGroup
	if err := s.scalarPostgresClient.Db.Preload("Custodials").Find(&custodialGroups).Error; err != nil {
		return nil, err
	}
	return custodialGroups, nil
}

func (s *ScalarClient) GetCustodialGroupByName(name string) (*models.CustodialGroup, error) {
	var custodialGroup models.CustodialGroup
	if err := s.scalarPostgresClient.Db.Where("name = ?", name).First(&custodialGroup).Error; err != nil {
		return nil, err
	}
	return &custodialGroup, nil
}

func (s *ScalarClient) GetCustodialGroupByID(id uint) (*models.CustodialGroup, error) {
	var custodialGroup models.CustodialGroup
	if err := s.scalarPostgresClient.Db.Preload("Custodials").Where("id = ?", id).First(&custodialGroup).Error; err != nil {
		return nil, err
	}
	return &custodialGroup, nil
}

func (s *ScalarClient) GetShortenCustodialGroups() ([]*models.ShortenCustodialGroup, error) {
	var shortenCustodialGroups []*models.ShortenCustodialGroup
	if err := s.scalarPostgresClient.Db.Model(&models.CustodialGroup{}).Select("id", "name", "btc_address").Find(&shortenCustodialGroups).Error; err != nil {
		return nil, err
	}
	return shortenCustodialGroups, nil
}
