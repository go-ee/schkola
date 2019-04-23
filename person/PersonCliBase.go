package person


type ChurchCli struct {
}

func NewChurchCli() (ret *ChurchCli) {
    ret = &ChurchCli{}
    return
}


type GraduationCli struct {
}

func NewGraduationCli() (ret *GraduationCli) {
    ret = &GraduationCli{}
    return
}


type ProfileCli struct {
}

func NewProfileCli() (ret *ProfileCli) {
    ret = &ProfileCli{}
    return
}


type PersonCli struct {
    ChurchCli *ChurchCli `json:"churchCli" eh:"optional"`
    GraduationCli *GraduationCli `json:"graduationCli" eh:"optional"`
    ProfileCli *ProfileCli `json:"profileCli" eh:"optional"`
}

func NewPersonCli() (ret *PersonCli) {
        
    churchCli := NewChurchCli()
    graduationCli := NewGraduationCli()
    profileCli := NewProfileCli()
    ret = &PersonCli{
        ChurchCli: churchCli,
        GraduationCli: graduationCli,
        ProfileCli: profileCli,
    }
    return
}









