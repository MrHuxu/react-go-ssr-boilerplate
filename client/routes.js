import Home from './components/home';
import Tags from './components/tags';

export default [
  { exact: true, path: '/', component: Home },
  { path: '/tags', component: Tags }
];
