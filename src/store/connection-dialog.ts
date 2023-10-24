import { create } from "zustand";

type ConnectionDialogState = {
    source: string | null;
    setSource: (source: string | null) => void;
};

export const useConnectionDialog = create<ConnectionDialogState>((set) => ({
    source: null,
    setSource: (source) => set({ source }),
}));