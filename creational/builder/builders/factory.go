package builders

func GetBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return &NormalBuilder{}
	}
	if builderType == "igloo" {
		return &IglooBuilder{}
	}
	return nil
}
