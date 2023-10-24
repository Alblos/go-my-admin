import { create } from "zustand";

type NavbarToggleState = {
    wide: boolean;
    toggleWide: () => void;
};

export const useNavbarToggle = create<NavbarToggleState>((set) => ({
    wide: true,
    toggleWide: () => set((state) => ({ wide: !state.wide })),
}));