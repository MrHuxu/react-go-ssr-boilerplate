import React from 'react';
import { renderToString } from 'react-dom/server';

import { Switch, Route, StaticRouter } from 'react-router';
import routes from './routes';

export const genHtmlString = url => renderToString(
  <StaticRouter url={ url }>
    <Switch>
      { routes.map(route => (
        <Route { ...route } />
      )) }
    </Switch>
  </StaticRouter>
);
