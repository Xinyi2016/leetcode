import dominantIndex from "./dominantIndex";

test("747. Largest Number At Least Twice of Others", () => {
  expect(dominantIndex([3,6,1,0])).toBe(1);
  expect(dominantIndex([1,2,3,4])).toBe(-1);
});
