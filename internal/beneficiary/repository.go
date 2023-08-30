package beneficiary

func (b BeneficiaryStruct) get_beneficiaries() ([]Beneficiary, error) {
	b.Logger.Println("Getting beneficiaries...")

	rows, err := b.DB.Query("SELECT id, name, address, email, currency, added_date, last_modified_date FROM beneficiaries")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beneficiaries []Beneficiary

	for rows.Next() {
		var beneficiary Beneficiary
		err := rows.Scan(&beneficiary.ID, &beneficiary.Name, &beneficiary.Address, &beneficiary.Email, &beneficiary.Currency, &beneficiary.AddedDate, &beneficiary.LastModifiedDate)
		if err != nil {
			return nil, err
		}
		beneficiaries = append(beneficiaries, beneficiary)
	}

	return beneficiaries, nil
}
