import React from 'react';
import classnames from 'classnames';

import './index.scss';

type BookmarkCardIconProps = {
  className?: string;
  black?: boolean;
  size: 'sm' | 'md' | 'lg';
};

const BookmarkCardIcon = ({
  size,
  black,
  className
}: BookmarkCardIconProps) => {
  const classes = classnames(
    'fa',
    'fa-bookmark',
    'bookmark__icon',
    {
      'fa-1x': size === 'sm'
    },
    {
      'fa-2x': size === 'md'
    },
    {
      'fa-3x': size === 'lg'
    },
    {
      black
    },
    className
  );

  return (
    <i
      className={classes}
      data-testid="system-health-icon"
      aria-hidden="true"
    />
  );
};

export default BookmarkCardIcon;
