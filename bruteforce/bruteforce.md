***Полный перебор (или метод «грубой силы»)***

Алгоритм построен на переборе. Последовательно сравниваются все символы текста между 0 и N - M, 
до тех пор пока не будет найдено первое совпадение с шаблоном или конец строки.
Далее, после каждой попытки, смещает символ шаблона ровно на один символ вправо.

The brute force algorithm requires no preprocessing phase, and a constant extra space in addition to the pattern and the
text. During the searching phase the text character comparisons can be done in any order. The time complexity of this
searching phase is O(mn) (when searching for am-1b in an for instance). The expected number of text character
comparisons is 2n.