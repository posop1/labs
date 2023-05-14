const getLength = (r: number): number | null => {
  if (!r || r === 0) {
    console.log('Числа нет или оно равно 0')

    return null
  }

  const p: number = 3.14
  const length = 2 * p * r

  console.log(length)

  return length
}

getLength(1)
