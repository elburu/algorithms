"use strict";

function merge(sortedLeftArr, sortedRightArr) {
  const resultArr = [];
  let leftIndex = 0;
  let rightIndex = 0;

  while (
    leftIndex < sortedLeftArr.length &&
    rightIndex < sortedRightArr.length
  ) {
    if (sortedLeftArr[leftIndex] <= sortedRightArr[rightIndex]) {
      resultArr.push(sortedLeftArr[leftIndex]);
      leftIndex++;
    } else {
      resultArr.push(sortedRightArr[rightIndex]);
      rightIndex++;
    }
  }
  while (leftIndex < sortedLeftArr.length) {
    resultArr.push(sortedLeftArr[leftIndex]);
    leftIndex++;
  }
  while (rightIndex < sortedRightArr.length) {
    resultArr.push(sortedRightArr[rightIndex]);
    rightIndex++;
  }

  return resultArr;
}

export function mergeSort(arr) {
  if (arr.length > 1) {
    const middle = arr.length / 2;

    const leftArr = arr.slice(0, middle);
    const rightArr = arr.slice(middle, arr.length);

    const sortedLeft = mergeSort(leftArr);
    const sortedRight = mergeSort(rightArr);

    return merge(sortedLeft, sortedRight);
  }
  return arr;
}
