import { getSnapshotChangedRequired, getSnapshotCurrentAllocation, getSnapshotTargetAllocation, getSnapshotTotalSummary } from "../advanced/calculations";

// Maturation dates in the tests are usually offset by a number of months and years using the current date and time. 
// I don't feel like writing date conversion logic so I'll be using static values here.
// It's easy and just as accurate to make static times using the same offset.

// Equivalent to the assets provided under `createsnapshot.go`.
// Calculations for a quick sanity check when computing group by values.
const assets: { holdingName: string, maturationDate?: string, total: number, taxShelter: string, assetCategory: string, institution: string, accountName: string, holdingER: number, skip: boolean }[] = [
    // Fidelity - Taxable
    { holdingName: "Holding1", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Fidelity", assetCategory: "TSM", accountName: "Account1", total: 250.25, holdingER: 0.3, skip: false },
    { holdingName: "Holding2", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Fidelity", assetCategory: "ITB", accountName: "Account1", total: 500.50, holdingER: 0.1, skip: false },

    // Fidelity - Roth
    { holdingName: "Holding1", maturationDate: undefined, taxShelter: "ROTH", institution: "Fidelity", assetCategory: "TSM", accountName: "Account2", total: 750.75, holdingER: 0.3, skip: false },
    { holdingName: "Holding2", maturationDate: undefined, taxShelter: "ROTH", institution: "Fidelity", assetCategory: "ITB", accountName: "Account2", total: 1000.00, holdingER: 0.1, skip: false },
];

const benchmark: { category: string, percent: number }[] = [
    { category: "TSM", percent: 60 },
    { category: "ITB", percent: 40 },
];

// Back of the envelope calculations for computing the different `Advanced Snapshot` test values
console.log(getSnapshotTotalSummary(assets));

console.log("Rebalance Figures...")
console.log(getSnapshotCurrentAllocation(assets));
console.log(getSnapshotTargetAllocation(assets, benchmark));
console.log(getSnapshotChangedRequired(assets, benchmark));