import React from 'react';
import { renderToString } from 'react-dom/server';

import { Provider } from 'react-redux';
import { createStaticStore } from './store';

import { Switch, Route, StaticRouter } from 'react-router';
import routes from './routes';

export const genHtmlString = (url, data) => {
  return renderToString(
    <Provider store={ createStaticStore(data) }>
      <StaticRouter location={ url }>
        <Switch>
          { routes.map(route => (
            <Route { ...route } />
          )) }
        </Switch>
      </StaticRouter>
    </Provider>
  );
};
