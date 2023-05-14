const getArea = (r: number) => {
  if (!r || r === 0) {
    console.log('Числа нет или оно равно 0')

    return null
  }

  const p: number = 3.14
  const area = p * r ** 2

  return area
}

console.log(getArea(2))
