export const logout = () => {
  console.log('lgout');
  // delete all localStorage items
  localStorage.clear();
  // clear history and redirect to login page (vue)
  window.location.href = '/login';
};
