import { z } from "zod";

const Note = z.object({
    id: z.number(),
    title: z.string(),
    content: z.string().optional(),
    sharelink: z.string().optional(),
    user_id: z.number().optional(),
});

const NoteResponse = z.object({
    data: z.object({
        note: Note.optional(),
        notes: z.array(Note).optional(),
    }).optional(),
    errors: z.record(z.string(), z.array(z.string())).optional(),
});

export async function fetchNote(id: number) {
    const response = await fetch(`http://127.0.0.1:3000/notes/${id}`, {
        method: "GET",
        credentials: "include"
    }).then((res) => res.json());

    const result = NoteResponse.safeParse(response);
    if (!result.success) {
        console.log(result.error);
    } else {
        return result.data
    }
}
