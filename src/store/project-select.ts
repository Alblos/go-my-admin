import { create } from "zustand";

type ProjectSelectState = {
    project: string;
    setProject: (project: string) => void;
};

export const useProjectSelect = create<ProjectSelectState>((set) => ({
    project: "Project 1",
    setProject: (project) => set(() => ({ project })),
}));