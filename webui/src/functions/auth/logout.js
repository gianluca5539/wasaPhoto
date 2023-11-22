export const logout = () => {
  console.log('lgout');
  // delete all localStorage items
  localStorage.clear();
  // go to /login
  window.location.href = '/login';
};
