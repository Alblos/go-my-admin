import {
    Dialog,
    DialogContent,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"
import { PlusIcon } from "lucide-react"
import { Button } from "../ui/button"
import SelectSource from "./select-source"
import { useConnectionDialog } from "@/store/connection-dialog"
import { AnimatePresence } from "framer-motion"
import ManualInput from "./manual-input"


type Props = {}



export default function AddConnection({ }: Props) {
    const { source, setSource } = useConnectionDialog()
    return (
        <Dialog>
            <DialogTrigger>
                <Button variant={"main"} onClick={() => setSource(null)}>
                    <PlusIcon size={25} />
                    Add Connection
                </Button>
            </DialogTrigger>
            <AnimatePresence>
                <DialogContent className="ov overflow-x-hidden max-w-3xl">
                    <DialogTitle>Add Connection</DialogTitle>
                    {!source && <SelectSource />}
                    {source == "manual" && <ManualInput />}
                </DialogContent>
            </AnimatePresence>

        </Dialog>
    )
}