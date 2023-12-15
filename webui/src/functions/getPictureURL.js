export const getPictureURL = (id) => {
  if (id == null || id == undefined || id == '') {
    return __API_URL__ + '/images/defaultprofile';
  }
  if (id == 'snoopy') {
    return __API_URL__ + '/images/defaultsnoopy';
  }
  return __API_URL__ + '/images/' + id;
};
