// You are a developer for a university. Your current project is to develop a system for students to find courses they share with friends. The university has a system for querying courses students are enrolled in, returned as a list of (ID, course) pairs.

// Write a function that takes in a collection of (student ID number, course name) pairs and returns, for every pair of students, a collection of all courses they share.


// Sample Input:

// enrollments1 = [
//   ["58", "Linear Algebra"],
//   ["94", "Art History"],
//   ["94", "Operating Systems"],
//   ["17", "Software Design"],
//   ["58", "Mechanics"],
//   ["58", "Economics"],
//   ["17", "Linear Algebra"],
//   ["17", "Political Science"],
//   ["94", "Economics"],
//   ["25", "Economics"],
//   ["58", "Software Design"],
// ]
// enrollments1 = [
//   ["58", ["Linear Algebra", "Mechanics", "Economics", "Software Design"]],
//   ["17", "Software Design, Political Science, Linear Algebra"],
// ]
// Creating all possible pairs:
// Sort and loop to push to result


// Sample Output (pseudocode, in any order):

// find_pairs(enrollments1) =>
// {
//   "58,17": ["Software Design", "Linear Algebra"]
//   "58,94": ["Economics"]
//   "58,25": ["Economics"]
//   "94,25": ["Economics"]
//   "17,94": []
//   "17,25": []
// }



// Additional test cases:

// Sample Input:

// enrollments2 = [
//   ["0", "Advanced Mechanics"],
//   ["0", "Art History"],
//   ["1", "Course 1"],
//   ["1", "Course 2"],
//   ["2", "Computer Architecture"],
//   ["3", "Course 1"],
//   ["3", "Course 2"],
//   ["4", "Algorithms"]
// ]



// Sample output:

// find_pairs(enrollments2) =>
// {
//   "1,0":[]
//   "2,0":[]
//   "2,1":[]
//   "3,0":[]
//   "3,1":["Course 1", "Course 2"]
//   "3,2":[]
//   "4,0":[]
//   "4,1":[]
//   "4,2":[]
//   "4,3":[]
// } 

// Sample Input:
// enrollments3 = [
//   ["23", "Software Design"], 
//   ["3", "Advanced Mechanics"], 
//   ["2", "Art History"], 
//   ["33", "Another"],
// ]


// Sample output:

// find_pairs(enrollments3) =>
// {
//   "23,3": []
//   "23,2": []
//   "23,33":[]
//   "3,2":  []
//   "3,33": []
//   "2,33": []
// }

// All Test Cases:
// find_pairs(enrollments1)
// find_pairs(enrollments2)
// find_pairs(enrollments3)

// Complexity analysis variables:

// n: number of student,course pairs in the input
// s: number of students
// c: total number of courses being offered (note: The number of courses any student can take is bounded by a small constant)

"use strict";
const _ = require('lodash');

const enrollments1 = [
  ["58", "Linear Algebra"],
  ["94", "Art History"],
  ["94", "Operating Systems"],
  ["17", "Software Design"],
  ["58", "Mechanics"],
  ["58", "Economics"],
  ["17", "Linear Algebra"],
  ["17", "Political Science"],
  ["94", "Economics"],
  ["25", "Economics"],
  ["58", "Software Design"]
];

const enrollments2 = [
  ["0", "Advanced Mechanics"],
  ["0", "Art History"],
  ["1", "Course 1"],
  ["1", "Course 2"],
  ["2", "Computer Architecture"],
  ["3", "Course 1"],
  ["3", "Course 2"],
  ["4", "Algorithms"]
];

const enrollments3 = [
  ["23", "Software Design"], 
  ["3",  "Advanced Mechanics"], 
  ["2",  "Art History"], 
  ["33", "Another"]
];

// +
// enrollments1 = [
//   ["58", ["Linear Algebra", "Mechanics", "Economics", "Software Design"]],
//   ["17", "Software Design, Political Science, Linear Algebra"],
// ]
// +
// Creating all possible pairs:
// +
// Sort and loop to push to result

// Time: 2*O(n) ~ O(n)
const mergeValuesOfKeys = (arr) => {
  const map = {}
  
  arr.map(([key, value]) => {
    map[key] = [...(map[key] || []), value]
  })
  for (const [key, value] of Object.entries(map)) {
    map[key] = value.sort()
  }
  return map
}

// O(s)*O(1) + 2*O(s) ~ 3*O(s) ~ O(s)
// array of strings
const findAllPairs = (inputMap) => {
  const keys = Object.keys(inputMap)
  const keyPairs = []
  
  keys.forEach((key, index) => {
    const tempKeys = [...keys]
    tempKeys.splice(index, 1)
    const pairs = tempKeys.map(v => ([key, v]))
    keyPairs.push(...pairs)
  })
  
  // remove duplicates
  const set = new Set(
    keyPairs.map(pair => pair.sort().join("-"))
  )
  const result = Array.from(set).map(v => v.split("-"))
  return result
}

// O((s*(s-1) / 2))
const findPairs = (uniqueMap, pairs) => {
  const resultMap = {}
  
  pairs.forEach(([first, second]) => {
    const key = [first, second].join(",")
    const arr1 = uniqueMap[first]
    const arr2 = uniqueMap[second]
    const intersection = _.intersection(arr1, arr2)
    
    
    resultMap[key] = intersection
  })
  
  return resultMap
}

// TOTAL: O(s) + O(n) + O((s*(s-1) / 2)) ~~~ O(n) + O(s) + O(s^2)
// Space: ???

const part1 = mergeValuesOfKeys(enrollments3)
// console.log(part1)

const part2 = findAllPairs(part1)
// console.log(part2)

const part3 = findPairs(part1, part2)
console.log(part3)
