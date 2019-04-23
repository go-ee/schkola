package library


type BookCli struct {
}

func NewBookCli() (ret *BookCli) {
    ret = &BookCli{}
    return
}


type LibraryCli struct {
    BookCli *BookCli `json:"bookCli" eh:"optional"`
}

func NewLibraryCli() (ret *LibraryCli) {
        
    bookCli := NewBookCli()
    ret = &LibraryCli{
        BookCli: bookCli,
    }
    return
}









