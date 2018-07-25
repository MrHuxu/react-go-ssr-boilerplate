import React from 'react';
import { renderToString } from 'react-dom/server';

import { Switch, Route, StaticRouter } from 'react-router';
import routes from './routes';

export const genHtmlString = url => {
  return renderToString(
    <StaticRouter location={ url }>
      <Switch>
        { routes.map(route => (
          <Route { ...route } />
        )) }
      </Switch>
    </StaticRouter>
  );
};
