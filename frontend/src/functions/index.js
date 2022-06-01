// Function to split by camelCase and uppercase the first letter
export const splitCamelCaseAndUppercase = (str) => {
  const val = str.replace(/([A-Z])/g, " $1").trim();
  return val.charAt(0).toUpperCase() + val.slice(1);
};

// Function to determinate if the value is an object
export const isObject = (value) => {
  if (value === null) {
    return false;
  }
  return typeof value === "object";
};
