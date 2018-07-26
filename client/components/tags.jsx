import React from 'react';
import { shape, arrayOf, string, objectOf, number } from 'prop-types';
import { connect } from 'react-redux';

const Tags = ({ data }) => {
  const { list, times } = data;

  return (
    <div>
      { list.map(item => (
        <p style={ { fontSize: 16 + times * 2 } }>
          { item }
        </p>
      )) }
    </div>
  );
};

Tags.propTypes = {
  data : shape({
    list  : arrayOf(string),
    times : objectOf(number)
  })
};

const mapStateToProps = ({ tags }) => ({ data: tags });

export default connect(mapStateToProps)(Tags);
