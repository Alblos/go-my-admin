import React from "react";
import { cva, type VariantProps } from "class-variance-authority";

const button = cva(["button", "w-fit", "flex", "justify-between", "items-center", "rounded-lg"], {
    variants: {
        intent: {
            primary: [
                "bg-slate-950", "hover:bg-slate-900/90", "text-white"
            ],
            secondary: [
                "bg-gray-100", "hover:bg-gray-200", "text-black"
            ],
            outline: [
                "bg-transparent hover:bg-gray-100 border border-gray-200", "text-black"
            ]
        },
        size: {
            base: ["text-base", "py-1", "px-2"],
            wide: ["text-base", "py-1", "px-4"],
            tall: ["text-base", "py-2", "px-2"],
            icon: ["text-base", "py-1", "px-1"],
        },
    },
    compoundVariants: [{ intent: "primary", size: "base" }],
    defaultVariants: {
        intent: "primary",
        size: "base",
    },
});

export interface ButtonProps
    extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof button> { }

export const Button: React.FC<ButtonProps> = ({
    className,
    intent,
    size,
    ...props
}) => <button className={button({ intent, size, className })} {...props} />;
