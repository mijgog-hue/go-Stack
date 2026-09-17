package main

type Stack struct{
    items []rune
}


func (s *Stack) Push(v rune){
    s.items = append(s.items,v)   //Короче добавляет файл
}


func (s *Stack) Pop()(rune,bool) {
if len(s.items) == 0{      // Проверка на пустоту Стака
    return 0,false
}


n := len(s.items) - 1  // узнает индекс последнего элемента
v := s.items[n]   // приравнивает v  к последнему элементу
s.items = s.items[:n] // Вот теперь она срезает слайс
return v,true  // и возвращаем последний элемент ВСЕ заебись👍 
}


func (s *Stack) Peek() (rune,bool){
     
    if len(s.items) == 0{         // Проверка на пустоту Стака
        return 0,false
    }
    return s.items[len(s.items)-1], true  // ПРосто смотрит последний элемент
}


func (s *Stack) IsEmpty() bool{
 return len(s.items) == 0  // Проверка на пустоту Стака
}

func (_ Solution) BalancedBrackets(s string) bool {
 stack := &Stack{} 
 pairs := map[rune]rune{
    ')': '(',
    ']': '[',
    '}': '{',
}

for _, char := range s{
	switch char {
    case '(', '[', '{':
        stack.Push(char)
    case ')', ']', '}':
       top, ok := stack.Pop()
				if ok == false || top!= pairs[char] {
			return false
		}
    }
}
return stack.IsEmpty()
}








Muvozanatlangan qavslar
Sizga turli qavslardan iborat va boshqa belgilar bo‘lishi mumkin bo‘lgan string beriladi. Qavslar muvozanatlanganmi yoki yo‘qligini aniqlovchi funksiya yozing.

Muvozanatlangan bo‘lishi uchun har bir ochilgan qavs to‘g‘ri yopilishi, tartib buzilmasligi va qavslar ichma-ich joylashishi shart.

Misol uchun
string = "([])(){}(())()()"
Kutilgan natija
true // bu qavslar muvozanatlangan
1-hint
2-hint

