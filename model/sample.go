package model

import "github.com/weet/service"

// sampleTableを作成する
func CreateSample(sample Sample) (Sample, error) {
	err := db.Create(&sample).Error
	if err != nil {
		return Sample{}, err
	}
	return sample, nil
}

func GetSampleById(sampleId uint) (service.Sample, error) {
	sample := service.Sample{}
	err := db.Find(&sample, sampleId).Error
	return sample, err
}
