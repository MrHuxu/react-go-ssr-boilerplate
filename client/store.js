export const createStaticStore = data => ({
  getState : () => JSON.parse(data)
});
