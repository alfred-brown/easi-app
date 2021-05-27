import React from 'react';
import { Link } from '@trussworks/react-uswds';
import classnames from 'classnames';

type LinkCardProps = {
  children: React.ReactNode;
  className?: string;
  link: string;
  heading: React.ReactNode | string;
} & JSX.IntrinsicElements['div'];

const LinkCard = ({
  children,
  className,
  link,
  heading,
  ...props
}: LinkCardProps) => {
  return (
    <div
      className={classnames(
        'padding-2',
        'line-height-body-4',
        'link-card-container',
        'bg-base-lightest',
        className
      )}
      {...props}
    >
      <h2 className="margin-top-0 margin-bottom-1">
        <Link href={link}>{heading}</Link>
      </h2>
      <div className="margin-top-1">{children}</div>
    </div>
  );
};

export default LinkCard;