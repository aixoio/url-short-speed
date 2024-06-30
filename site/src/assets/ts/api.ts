
export interface Short {
    short: string
}

export async function MakeShort(long: string): Promise<Short> {
    const result = await fetch("/api/link/new", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            long
        })
    })

    const data: Short = await result.json();

    return data
}
