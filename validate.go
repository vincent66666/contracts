package contracts

// Checker 检查器
type Checker interface {
	Check(value interface{}) error
}

// Checkers 检查器组
type Checkers map[string][]Checker

// ValidateErrors 验证错误信息
type ValidateErrors map[string][]string

// ValidatedResult 验证结果
type ValidatedResult interface {
	SafeValidate
	IsFail() bool
	IsSuccessful() bool
	Errors() ValidateErrors
}

// FieldsAlias 有别名
type FieldsAlias interface {
	GetFieldsNameMap() map[string]string
}

// Validator 验证器
type Validator interface {
	Validate() ValidatedResult
}

// SafeValidate 验证不通过即 panic
type SafeValidate interface {
	Assure()
}

// ValidatableForm 可验证的表单
type ValidatableForm interface {
	GetCheckers() Checkers
	GetFields() Fields
}

// ValidatableAliasForm 可验证并且设置了字段名映射的表单
type ValidatableAliasForm interface {
	FieldsAlias
	GetCheckers() Checkers
	GetFields() Fields
}
