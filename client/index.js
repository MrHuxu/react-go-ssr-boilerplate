import React from 'react';
import { renderToString } from 'react-dom/server';

const App = () => (
  <div>
    <h1> Hello SSR! </h1>
  </div>
);

export const html = renderToString(<App />);
