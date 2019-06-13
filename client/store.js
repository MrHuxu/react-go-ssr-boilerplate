export const createStaticStore = data => ({
  getState : data ? () => data : {}
});
