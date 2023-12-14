export const getPictureURL = (id) => {
  if (id == null || id == undefined || id == '') {
    return __API_URL__ + '/images/default';
  }
  return __API_URL__ + '/images/' + id;
};
