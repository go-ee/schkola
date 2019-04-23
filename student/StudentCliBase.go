package student


type AttendanceCli struct {
}

func NewAttendanceCli() (ret *AttendanceCli) {
    ret = &AttendanceCli{}
    return
}


type CourseCli struct {
}

func NewCourseCli() (ret *CourseCli) {
    ret = &CourseCli{}
    return
}


type GradeCli struct {
}

func NewGradeCli() (ret *GradeCli) {
    ret = &GradeCli{}
    return
}


type GroupCli struct {
}

func NewGroupCli() (ret *GroupCli) {
    ret = &GroupCli{}
    return
}


type SchoolApplicationCli struct {
}

func NewSchoolApplicationCli() (ret *SchoolApplicationCli) {
    ret = &SchoolApplicationCli{}
    return
}


type SchoolYearCli struct {
}

func NewSchoolYearCli() (ret *SchoolYearCli) {
    ret = &SchoolYearCli{}
    return
}


type StudentCli struct {
    AttendanceCli *AttendanceCli `json:"attendanceCli" eh:"optional"`
    CourseCli *CourseCli `json:"courseCli" eh:"optional"`
    GradeCli *GradeCli `json:"gradeCli" eh:"optional"`
    GroupCli *GroupCli `json:"groupCli" eh:"optional"`
    SchoolApplicationCli *SchoolApplicationCli `json:"schoolApplicationCli" eh:"optional"`
    SchoolYearCli *SchoolYearCli `json:"schoolYearCli" eh:"optional"`
}

func NewStudentCli() (ret *StudentCli) {
        
    attendanceCli := NewAttendanceCli()
    courseCli := NewCourseCli()
    gradeCli := NewGradeCli()
    groupCli := NewGroupCli()
    schoolApplicationCli := NewSchoolApplicationCli()
    schoolYearCli := NewSchoolYearCli()
    ret = &StudentCli{
        AttendanceCli: attendanceCli,
        CourseCli: courseCli,
        GradeCli: gradeCli,
        GroupCli: groupCli,
        SchoolApplicationCli: schoolApplicationCli,
        SchoolYearCli: schoolYearCli,
    }
    return
}









