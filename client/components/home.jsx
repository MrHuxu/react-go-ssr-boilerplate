import React from 'react';
import { shape, arrayOf, string, object, objectOf, number } from 'prop-types';
import { connect } from 'react-redux';

const Home = ({ data }) => {
  const { titles, infos } = data;

  return (
    <div>
      <p> Life of xhu >> </p>
      { titles.map(title => (
        <div>
          <p> { infos[title].seq } </p>
          <p> { infos[title].title } </p>
          <p> { infos[title].time } </p>
          <p> { infos[title].tags.join(',') } </p>
          <p> { infos[title].content.slice(0, 100) } </p>
        </div>
      )) }
    </div>
  );
};

Home.propTypes = {
  data : shape({
    titles : arrayOf(string),
    infos  : objectOf(shape({
      seq     : number,
      title   : string,
      time    : object,
      tags    : arrayOf(string),
      content : string
    }))
  })
};

const mapStateToProps = ({ home }) => ({ home });

export default connect(mapStateToProps)(Home);
