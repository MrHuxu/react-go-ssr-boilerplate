import React from 'react';
import { shape, arrayOf, string, objectOf, number } from 'prop-types';
import { connect } from 'react-redux';

const Tags = ({ data }) => {
  const { list, infos } = data;

  return (
    <div>
      { list.map(item => (
        <p style={ { fontSize: 16 + infos[item] * 10 } }>
          { item }
        </p>
      )) }
      <a href="/"> back to home </a>
    </div>
  );
};

Tags.propTypes = {
  data : shape({
    list  : arrayOf(string),
    infos : objectOf(number)
  })
};

const mapStateToProps = ({ test }) => ({ data: test });

export default connect(mapStateToProps)(Tags);
