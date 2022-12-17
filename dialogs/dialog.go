package dialogs

import (
	"sync"
)

var openedDialogs = make(map[int64]*Dialog, 0)

var mutex sync.Mutex

type Dialog struct {
	UserID    int64
	UserName  string
	FirstName string
	LastName  string
	ChatID    int64
	Reply     string
	Replied   bool
}

func GetDialog(userID int64) *Dialog {
	mutex.Lock()
	dialog, ok := openedDialogs[userID]
	mutex.Unlock()

	if ok {
		return dialog
	} else {
		return nil
	}
}

func SaveDialog(dialog *Dialog) {
	mutex.Lock()
	openedDialogs[dialog.UserID] = dialog
	mutex.Unlock()
}

func RemoveDialog(userID int64) {
	mutex.Lock()
	openedDialogs[userID] = nil
	mutex.Unlock()
}
