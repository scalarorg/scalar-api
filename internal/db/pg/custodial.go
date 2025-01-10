package pg

import "github.com/scalarorg/xchains-api/internal/db/pg/models"

func (c *PostgresClient) SaveCustodial(custodial *models.Custodial) error {
	return c.DB.Create(custodial).Error
}

func (c *PostgresClient) GetCustodials() ([]*models.Custodial, error) {
	var custodials []*models.Custodial
	if err := c.DB.Find(&custodials).Error; err != nil {
		return nil, err
	}
	return custodials, nil
}

func (c *PostgresClient) GetCustodialByName(name string) (*models.Custodial, error) {
	var custodial models.Custodial
	if err := c.DB.Where("name = ?", name).First(&custodial).Error; err != nil {
		return nil, err
	}
	return &custodial, nil
}

func (c *PostgresClient) GetCustodialByNames(names []string) ([]*models.Custodial, error) {
	var custodials []*models.Custodial
	if err := c.DB.Where("name IN (?)", names).Find(&custodials).Error; err != nil {
		return nil, err
	}
	return custodials, nil
}

func (c *PostgresClient) SaveCustodialGroup(custodialGroup *models.CustodialGroup) error {
	return c.DB.Create(custodialGroup).Error
}

func (c *PostgresClient) GetCustodialGroups() ([]*models.CustodialGroup, error) {
	var custodialGroups []*models.CustodialGroup
	if err := c.DB.Preload("Custodials").Find(&custodialGroups).Error; err != nil {
		return nil, err
	}
	return custodialGroups, nil
}

func (c *PostgresClient) GetCustodialGroupByName(name string) (*models.CustodialGroup, error) {
	var custodialGroup models.CustodialGroup
	if err := c.DB.Where("name = ?", name).First(&custodialGroup).Error; err != nil {
		return nil, err
	}
	return &custodialGroup, nil
}

func (c *PostgresClient) GetCustodialGroupByID(id uint) (*models.CustodialGroup, error) {
	var custodialGroup models.CustodialGroup
	if err := c.DB.Preload("Custodials").Where("id = ?", id).First(&custodialGroup).Error; err != nil {
		return nil, err
	}
	return &custodialGroup, nil
}

func (c *PostgresClient) GetShortenCustodialGroups() ([]*models.ShortenCustodialGroup, error) {
	var shortenCustodialGroups []*models.ShortenCustodialGroup
	if err := c.DB.Model(&models.CustodialGroup{}).Select("id", "name", "taproot_address").Find(&shortenCustodialGroups).Error; err != nil {
		return nil, err
	}
	return shortenCustodialGroups, nil
}
