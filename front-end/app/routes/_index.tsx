import type { MetaFunction } from "@remix-run/node";

import type { LoaderFunctionArgs  } from "@remix-run/node";
import { json } from "@remix-run/node";
import { fs } from "~/utils/fs-promises.server";
import { useLoaderData } from '@remix-run/react'

// Define TypeScript type for ReviewRecord
type ReviewRecord = {
	content: string;
	author: string;
	score: string;
	timestamp: string;
	id: string;
};

export const loader = async (args: LoaderFunctionArgs ) => {
  // Find the absolute path of the json directory
  const jsonDirectory = process.cwd() + "/json";
  // Read the json data file data.json
  const fileContents = await fs.readFile(jsonDirectory + "/data.json", "utf8");
  // Parse the json data file contents into a json object
  const data: ReviewRecord[] = JSON.parse(fileContents);

  return json({
    data,
  });
};

export default function Index() {
	const { data } = useLoaderData<{ data: ReviewRecord[] }>();

	return (
		<div>
			<h1>apple-app-store-reviews-service</h1>
			<table>
                <tr>
                    <th>Timestamp</th>
                    <th>Author</th>
					<th>Score</th>
                    <th>Content</th>
                </tr>
                {data.map((record) => (
				<tr>
              		<td>{record.timestamp}</td>
              		<td>{record.author}</td>
              		<td>{record.score}</td>
              		<td>{record.content}</td>
            	</tr>
          ))}
            </table>
		</div>
	)
}
