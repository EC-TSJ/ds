/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package ds_test

import (
	dal "ds"
	"fmt"
	"testing"
)

func TestFuncs(t *testing.T) {
	qt := dal.NewQueue()
	qt.Enqueue("uno", "dos", 5, "tres", "cuatro", "cinco", "seis")
	fmt.Println(qt)
	var ss interface{} = qt.Dequeue()
	if ss == nil {
		return
	}
	var ff interface{} = qt.Peek()
	if ff == nil {
		return
	}
	fmt.Println(qt)
	var tt int = qt.Length()
	var dd bool = qt.Empty()
	if tt == 26 && dd {
		return
	}
	qt.Remove()

	fmt.Println(qt)
	qt.Enqueue("uno", "dos", 5, "tres", "cuatro", "cinco")

	fmt.Println(qt)

}
