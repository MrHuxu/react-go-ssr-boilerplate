import React from 'react';
import { shape, arrayOf, string, objectOf, number } from 'prop-types';
import { connect } from 'react-redux';

const Home = ({ data }) => {
  const { titles, infos } = data;

  return (
    <div>
      <p> Life of xhu >> </p>
      { titles.map(title => (
        <div>
          <p> { title } </p>
          <p> { infos[title] } </p>
          <a href="/test"> to test </a>
        </div>
      )) }
    </div>
  );
};

Home.propTypes = {
  data : shape({
    list  : arrayOf(string),
    infos : objectOf(number)
  })
};

const mapStateToProps = ({ home }) => ({ data: home });

export default connect(mapStateToProps)(Home);
