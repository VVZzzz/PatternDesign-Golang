// Package Observable 即被观察/订阅者,也可称之为 subject
package observe

import (
	"errors"
	"fmt"
)

type Observable struct {
	changed      bool
	observerList []Observer
}

func (ob *Observable) RegisterObservers(observer Observer) error {
	if len(ob.observerList) <= 0 {
		return errors.New("observerList invalid")
	}
	if ob.observerList[len(ob.observerList)-1] == observer {
		return errors.New("observer to add has exist")
	}
	ob.observerList = append(ob.observerList, observer)
	return nil
}

func (ob *Observable) RemoveObservers(observer Observer) error {
	for idx, v := range ob.observerList {
		if v == observer {
			ob.observerList = append(ob.observerList[:idx], ob.observerList[idx+1:]...)
			return nil
		}
	}
	return errors.New("not found observer in list")
}

// NotifyObservers 该方法签名中有 data 时,实际是一个 push 的实现,将特定数据给 push 到各观察者上
func (ob *Observable) NotifyObservers(observer Observer, data interface{}) {
	if !ob.changed {
		fmt.Println("no change happened.")
		return
	}
	for _, o := range ob.observerList {
		o.Update(ob, data)
	}
	ob.changed = false
}
