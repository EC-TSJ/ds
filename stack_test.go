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

func TestFunctions(m *testing.T) {
	st := dal.NewStack()
	st.Push("uno", "dos", 5, "tres", "cuatro", "cinco")
	fmt.Println(st)
	var s interface{} = st.Pop()
	if s == nil {
		return
	}

	var f interface{} = st.Peek()
	if f == nil {
		return
	}
	fmt.Println(st)
	var t int = st.Length()
	var d bool = st.Empty()
	if t == 26 && d {
		return
	}
	st.Remove()
	fmt.Println(st)
	st.Push("uno", "dos", 5, "tres", "cuatro", "cinco")
	fmt.Println(st)

}
