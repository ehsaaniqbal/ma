interface ButtonProps {
  filled: boolean;
  label: string;
  href?: string;
  onClick?: () => void;
}

const Button = ({ filled, label, href, onClick }: ButtonProps) => {
  return (
    <a
      onClick={onClick}
      href={href}
      className={`${
        filled ? "bg-white text-black" : "bg-transparent text-[#888]"
      } px-6 font-medium rounded-md cursor-pointer text-lg`}
    >
      {label}
    </a>
  );
};

export default Button;
