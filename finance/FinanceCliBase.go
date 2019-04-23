package finance


type ExpenseCli struct {
}

func NewExpenseCli() (ret *ExpenseCli) {
    ret = &ExpenseCli{}
    return
}


type ExpensePurposeCli struct {
}

func NewExpensePurposeCli() (ret *ExpensePurposeCli) {
    ret = &ExpensePurposeCli{}
    return
}


type FeeCli struct {
}

func NewFeeCli() (ret *FeeCli) {
    ret = &FeeCli{}
    return
}


type FeeKindCli struct {
}

func NewFeeKindCli() (ret *FeeKindCli) {
    ret = &FeeKindCli{}
    return
}


type FinanceCli struct {
    ExpenseCli *ExpenseCli `json:"expenseCli" eh:"optional"`
    ExpensePurposeCli *ExpensePurposeCli `json:"expensePurposeCli" eh:"optional"`
    FeeCli *FeeCli `json:"feeCli" eh:"optional"`
    FeeKindCli *FeeKindCli `json:"feeKindCli" eh:"optional"`
}

func NewFinanceCli() (ret *FinanceCli) {
        
    expenseCli := NewExpenseCli()
    expensePurposeCli := NewExpensePurposeCli()
    feeCli := NewFeeCli()
    feeKindCli := NewFeeKindCli()
    ret = &FinanceCli{
        ExpenseCli: expenseCli,
        ExpensePurposeCli: expensePurposeCli,
        FeeCli: feeCli,
        FeeKindCli: feeKindCli,
    }
    return
}









