export const getPictureURL = (id) => {
  if (id == null || id == undefined || id == '') {
    return 'http://localhost:3000/images/default';
  }
  return 'http://localhost:3000/images/' + id;
};
