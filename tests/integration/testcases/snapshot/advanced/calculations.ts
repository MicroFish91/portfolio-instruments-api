// Maturation dates in the tests are usually offset by a number of months and years using the current date and time. 
// I don't feel like writing date conversion logic so I'll be using static values here.
// It's easy and just as accurate to make static times using the same offset.

const mockDateOfOrigin: string = "07/01/2000";

// Equivalent to the assets provided under `createsnapshot.go`.
// Calculations for a quick sanity check when computing group by values.
const advancedAssets: { holdingName: string, maturationDate?: string, total: number, taxShelter: string, assetCategory: string, institution: string, accountName: string, holdingER: number, skip: boolean }[] = [
    // Indices -
    // Accounts: 0
    // Holdings: 2, 12, 6, 3, 7, 5, 9
    { holdingName: "Holding3", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Vanguard", assetCategory: "CASH", accountName: "Account1", total: 10341.01, holdingER: 0, skip: false },
    { holdingName: "Holding13", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Vanguard", assetCategory: "OTHER", accountName: "Account1", total: 11979.70, holdingER: 1.18, skip: false },
    { holdingName: "Holding7", maturationDate: "01/01/2003", taxShelter: "TAXABLE", institution: "Vanguard", assetCategory: "STB", accountName: "Account1", total: 1092.52, holdingER: 0, skip: false },
    { holdingName: "Holding4", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Vanguard", assetCategory: "GOLD", accountName: "Account1", total: 3683.72, holdingER: 0.82, skip: false },
    { holdingName: "Holding8", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Vanguard", assetCategory: "TSM", accountName: "Account1", total: 1913.09, holdingER: 0.18, skip: false },
    { holdingName: "Holding6", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Vanguard", assetCategory: "OTHER", accountName: "Account1", total: 3891.44, holdingER: 0.58, skip: false },
    { holdingName: "Holding10", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Vanguard", assetCategory: "CASH", accountName: "Account1", total: 3626.10, holdingER: 0, skip: true },

    // Indices -
    // Accounts: 1
    // Holdings: 3, 1, 13, 10, 8, 4, 0
    { holdingName: "Holding4", maturationDate: undefined, taxShelter: "ROTH", institution: "Vanguard", assetCategory: "GOLD", accountName: "Account2", total: 9371.45, holdingER: 0.82, skip: false },
    { holdingName: "Holding2", maturationDate: "01/01/2029", taxShelter: "ROTH", institution: "Vanguard", assetCategory: "LTB", accountName: "Account2", total: 2271.85, holdingER: 0, skip: false },
    { holdingName: "Holding14", maturationDate: "07/01/2010", taxShelter: "ROTH", institution: "Vanguard", assetCategory: "ITB", accountName: "Account2", total: 12096.14, holdingER: 0, skip: false },
    { holdingName: "Holding11", maturationDate: undefined, taxShelter: "ROTH", institution: "Vanguard", assetCategory: "GOLD", accountName: "Account2", total: 7020.39, holdingER: 0.77, skip: false },
    { holdingName: "Holding9", maturationDate: "07/01/2026", taxShelter: "ROTH", institution: "Vanguard", assetCategory: "LTB", accountName: "Account2", total: 13431.37, holdingER: 0, skip: false },
    { holdingName: "Holding5", maturationDate: undefined, taxShelter: "ROTH", institution: "Vanguard", assetCategory: "DSCV", accountName: "Account2", total: 7728.62, holdingER: 0.64, skip: false },
    { holdingName: "Holding1", maturationDate: undefined, taxShelter: "ROTH", institution: "Vanguard", assetCategory: "TSM", accountName: "Account2", total: 11658.71, holdingER: 0.3, skip: true },

    // Indices -
    // Accounts: 2
    // Holdings: 4, 9, 11, 5, 2, 3, 6
    { holdingName: "Holding5", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Vanguard", assetCategory: "DSCV", accountName: "Account3", total: 7780.79, holdingER: 0.64, skip: false },
    { holdingName: "Holding10", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Vanguard", assetCategory: "CASH", accountName: "Account3", total: 6969.85, holdingER: 0, skip: false },
    { holdingName: "Holding12", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Vanguard", assetCategory: "DSCV", accountName: "Account3", total: 10443.17, holdingER: 0.9, skip: false },
    { holdingName: "Holding6", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Vanguard", assetCategory: "OTHER", accountName: "Account3", total: 12502.99, holdingER: 0.58, skip: false },
    { holdingName: "Holding3", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Vanguard", assetCategory: "CASH", accountName: "Account3", total: 286.65, holdingER: 0, skip: false },
    { holdingName: "Holding4", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Vanguard", assetCategory: "GOLD", accountName: "Account3", total: 15706.22, holdingER: 0.82, skip: false },
    { holdingName: "Holding7", maturationDate: "01/01/2003", taxShelter: "TRADITIONAL", institution: "Vanguard", assetCategory: "STB", accountName: "Account3", total: 5880.80, holdingER: 0, skip: true },

    // Indices -
    // Accounts: 3
    // Holdings: 3, 7, 4, 9, 0, 2, 11
    { holdingName: "Holding4", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Fidelity", assetCategory: "GOLD", accountName: "Account4", total: 7625.72, holdingER: 0.82, skip: false },
    { holdingName: "Holding8", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Fidelity", assetCategory: "TSM", accountName: "Account4", total: 10586.65, holdingER: 0.18, skip: false },
    { holdingName: "Holding5", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Fidelity", assetCategory: "DSCV", accountName: "Account4", total: 7324.53, holdingER: 0.64, skip: false },
    { holdingName: "Holding10", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Fidelity", assetCategory: "CASH", accountName: "Account4", total: 1168.83, holdingER: 0, skip: false },
    { holdingName: "Holding1", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Fidelity", assetCategory: "TSM", accountName: "Account4", total: 5268.47, holdingER: 0.3, skip: false },
    { holdingName: "Holding3", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Fidelity", assetCategory: "CASH", accountName: "Account4", total: 4929.89, holdingER: 0, skip: false },
    { holdingName: "Holding12", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Fidelity", assetCategory: "DSCV", accountName: "Account4", total: 437.78, holdingER: 0.9, skip: true },

    // Indices -
    // Accounts: 4
    // Holdings: 8, 9, 4, 12, 10, 6, 4
    { holdingName: "Holding9", maturationDate: "07/01/2026", taxShelter: "ROTH", institution: "Fidelity", assetCategory: "LTB", accountName: "Account5", total: 15407.15, holdingER: 0, skip: false },
    { holdingName: "Holding10", maturationDate: undefined, taxShelter: "ROTH", institution: "Fidelity", assetCategory: "CASH", accountName: "Account5", total: 4662.95, holdingER: 0, skip: false },
    { holdingName: "Holding5", maturationDate: undefined, taxShelter: "ROTH", institution: "Fidelity", assetCategory: "DSCV", accountName: "Account5", total: 13466.88, holdingER: 0.64, skip: false },
    { holdingName: "Holding13", maturationDate: undefined, taxShelter: "ROTH", institution: "Fidelity", assetCategory: "OTHER", accountName: "Account5", total: 4747.47, holdingER: 1.18, skip: false },
    { holdingName: "Holding11", maturationDate: undefined, taxShelter: "ROTH", institution: "Fidelity", assetCategory: "GOLD", accountName: "Account5", total: 515.56, holdingER: 0.77, skip: false },
    { holdingName: "Holding7", maturationDate: "01/01/2003", taxShelter: "ROTH", institution: "Fidelity", assetCategory: "STB", accountName: "Account5", total: 6889.24, holdingER: 0, skip: false },
    { holdingName: "Holding5", maturationDate: undefined, taxShelter: "ROTH", institution: "Fidelity", assetCategory: "DSCV", accountName: "Account5", total: 13695.74, holdingER: 0.64, skip: true },

    // Indices -
    // Accounts: 5
    // Holdings: 4, 12, 2, 6, 1, 7, 13
    { holdingName: "Holding5", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Fidelity", assetCategory: "DSCV", accountName: "Account6", total: 12673.13, holdingER: 0.64, skip: false },
    { holdingName: "Holding13", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Fidelity", assetCategory: "OTHER", accountName: "Account6", total: 3073.30, holdingER: 1.18, skip: false },
    { holdingName: "Holding3", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Fidelity", assetCategory: "CASH", accountName: "Account6", total: 11547.25, holdingER: 0, skip: false },
    { holdingName: "Holding7", maturationDate: "01/01/2003", taxShelter: "TRADITIONAL", institution: "Fidelity", assetCategory: "STB", accountName: "Account6", total: 2240.44, holdingER: 0, skip: false },
    { holdingName: "Holding2", maturationDate: "01/01/2029", taxShelter: "TRADITIONAL", institution: "Fidelity", assetCategory: "LTB", accountName: "Account6", total: 3483.74, holdingER: 0, skip: false },
    { holdingName: "Holding8", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Fidelity", assetCategory: "TSM", accountName: "Account6", total: 1858.15, holdingER: 0.18, skip: false },
    { holdingName: "Holding14", maturationDate: "07/01/2010", taxShelter: "TRADITIONAL", institution: "Fidelity", assetCategory: "ITB", accountName: "Account6", total: 9448.52, holdingER: 0, skip: true },

    // Indices -
    // Accounts: 6
    // Holdings: 1, 13, 11, 3, 5, 12, 10
    { holdingName: "Holding2", maturationDate: "01/01/2029", taxShelter: "TAXABLE", institution: "Schwab", assetCategory: "LTB", accountName: "Account7", total: 15130.08, holdingER: 0, skip: false },
    { holdingName: "Holding14", maturationDate: "07/01/2010", taxShelter: "TAXABLE", institution: "Schwab", assetCategory: "ITB", accountName: "Account7", total: 11453.51, holdingER: 0, skip: false },
    { holdingName: "Holding12", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Schwab", assetCategory: "DSCV", accountName: "Account7", total: 7429.49, holdingER: 0.9, skip: false },
    { holdingName: "Holding4", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Schwab", assetCategory: "GOLD", accountName: "Account7", total: 1639.83, holdingER: 0.82, skip: false },
    { holdingName: "Holding6", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Schwab", assetCategory: "OTHER", accountName: "Account7", total: 14835.84, holdingER: 0.58, skip: false },
    { holdingName: "Holding13", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Schwab", assetCategory: "OTHER", accountName: "Account7", total: 9678.95, holdingER: 1.18, skip: false },
    { holdingName: "Holding11", maturationDate: undefined, taxShelter: "TAXABLE", institution: "Schwab", assetCategory: "GOLD", accountName: "Account7", total: 10854.47, holdingER: 0.77, skip: true },

    // Indices -
    // Accounts: 7
    // Holdings: 6, 12, 5, 13, 2, 11, 4
    { holdingName: "Holding7", maturationDate: "01/01/2003", taxShelter: "ROTH", institution: "Schwab", assetCategory: "STB", accountName: "Account8", total: 15737.92, holdingER: 0, skip: false },
    { holdingName: "Holding13", maturationDate: undefined, taxShelter: "ROTH", institution: "Schwab", assetCategory: "OTHER", accountName: "Account8", total: 9434.05, holdingER: 1.18, skip: false },
    { holdingName: "Holding6", maturationDate: undefined, taxShelter: "ROTH", institution: "Schwab", assetCategory: "OTHER", accountName: "Account8", total: 10096.92, holdingER: 0.58, skip: false },
    { holdingName: "Holding14", maturationDate: "07/01/2010", taxShelter: "ROTH", institution: "Schwab", assetCategory: "ITB", accountName: "Account8", total: 15411.10, holdingER: 0, skip: false },
    { holdingName: "Holding3", maturationDate: undefined, taxShelter: "ROTH", institution: "Schwab", assetCategory: "CASH", accountName: "Account8", total: 6899.97, holdingER: 0, skip: false },
    { holdingName: "Holding12", maturationDate: undefined, taxShelter: "ROTH", institution: "Schwab", assetCategory: "DSCV", accountName: "Account8", total: 4343.31, holdingER: 0.9, skip: false },
    { holdingName: "Holding5", maturationDate: undefined, taxShelter: "ROTH", institution: "Schwab", assetCategory: "DSCV", accountName: "Account8", total: 4277.00, holdingER: 0.64, skip: true },

    // Indices -
    // Accounts: 8
    // Holdings: 2, 9, 6, 8, 5, 0, 13
    { holdingName: "Holding3", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Schwab", assetCategory: "CASH", accountName: "Account9", total: 14045.3, holdingER: 0, skip: false },
    { holdingName: "Holding10", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Schwab", assetCategory: "CASH", accountName: "Account9", total: 11642.63, holdingER: 0, skip: false },
    { holdingName: "Holding7", maturationDate: "01/01/2003", taxShelter: "TRADITIONAL", institution: "Schwab", assetCategory: "STB", accountName: "Account9", total: 6618.72, holdingER: 0, skip: false },
    { holdingName: "Holding9", maturationDate: "07/01/2026", taxShelter: "TRADITIONAL", institution: "Schwab", assetCategory: "LTB", accountName: "Account9", total: 1124.46, holdingER: 0, skip: false },
    { holdingName: "Holding6", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Schwab", assetCategory: "OTHER", accountName: "Account9", total: 13441.36, holdingER: 0.58, skip: false },
    { holdingName: "Holding1", maturationDate: undefined, taxShelter: "TRADITIONAL", institution: "Schwab", assetCategory: "TSM", accountName: "Account9", total: 10426.11, holdingER: 0.3, skip: false },
    { holdingName: "Holding14", maturationDate: "07/01/2010", taxShelter: "TRADITIONAL", institution: "Schwab", assetCategory: "ITB", accountName: "Account9", total: 11859.11, holdingER: 0, skip: true },
];

// Back of the envelope calculations for computing the different `Advanced Snapshot` test values
console.log(getSnapshotTotalSummary(advancedAssets));
console.log(parseSnapshotWithGroupByKey("accountName"));
console.log(parseSnapshotWithGroupByKey("assetCategory"));
console.log(parseSnapshotWithGroupByKey("institution"));
console.log(parseSnapshotWithGroupByKey("taxShelter"));

console.log("Maturation Dates Filtering...");
console.log(getSnapshotByMaturationDateRange(undefined, undefined));
console.log(getSnapshotByMaturationDateRange("01/01/2028", undefined));
console.log(getSnapshotByMaturationDateRange(undefined, "08/01/2011"));

export function getSnapshotTotalSummary(assets: typeof advancedAssets): { sum: number, er: number } {
    let sum = 0;
    let erSum = 0;
    for (const { total, holdingER } of assets) {
        sum += total;
        erSum += total * holdingER;
    }

    return { sum, er: erSum / sum };
}

export function parseSnapshotWithGroupByKey(groupByKey: keyof typeof advancedAssets[number]): { fields: string[], totals: number[] } {
    if (typeof groupByKey != 'string' && typeof groupByKey != 'number') {
        throw new Error("Incompatible groupByKey type");
    }

    const resourcesGrouped: Record<string | number, number> = {};
    for (const asset of advancedAssets) {
        const key = asset[groupByKey] as string | number;
        resourcesGrouped[key] = (resourcesGrouped[key] ?? 0) + asset.total;
    }

    const fields: string[] = [];
    const totals: number[] = [];
    for (const [key, val] of Object.entries(resourcesGrouped)) {
        fields.push(key);
        totals.push(val);
    }

    return { fields, totals };
}


export function getSnapshotByMaturationDateRange(maturationStart?: string, maturationEnd?: string) {
    console.log("Start Date: ", maturationStart, ", End Date: ", maturationEnd);
    maturationStart ??= mockDateOfOrigin;

    const resources: any = [];  // Don't care about strict typing, just use any
    for (const asset of advancedAssets) {
        if (!asset.maturationDate) {
            continue;
        }

        const startDate: Date = new Date(maturationStart);
        const endDate: Date | undefined = maturationEnd ? new Date(maturationEnd) : undefined;
        const maturationDate: Date = new Date(asset.maturationDate);

        if (maturationDate < startDate) {
            continue;
        }

        if (endDate && maturationDate > endDate) {
            continue;
        }

        resources.push({
            Account_name: asset.accountName,
            Holding_name: asset.holdingName,
            Asset_category: asset.assetCategory,
            Maturation_date: asset.maturationDate,
            Total: asset.total,
            Skip_rebalance: asset.skip,
        });
    }

    return resources;
}