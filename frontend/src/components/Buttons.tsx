export const Button = ({
  className,
  ...props
}: {
  className?: string;
} & React.ButtonHTMLAttributes<HTMLButtonElement>) => {
  return (
    <>
      <button className={className} {...props} />
    </>
  );
};
