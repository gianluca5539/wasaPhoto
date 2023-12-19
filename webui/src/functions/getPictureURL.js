export const getPictureURL = (id) => {
  if (id == null || id == undefined || id == -1) {
    return __API_URL__ + '/images/defaultprofile';
  }
  if (id == -2) {
    return __API_URL__ + '/images/defaultsnoopy';
  }
  return __API_URL__ + '/images/' + id;
};
