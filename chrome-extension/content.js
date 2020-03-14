const lc = location.href;
const player = "http://localhost:3001/fuinki?u=";
fetch(encodeURI(player + location.href), {
  mode: 'cors'
});
