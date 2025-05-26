#include <unordered_map>
#include <vector>
#include <iostream>

using namespace std;

class Solution
{
public:
    vector<int> twoSum(vector<int> &nums, int target)
    {
        std::unordered_map<int, int> numMap;

        for (int i = 0; i < nums.size(); ++i)
        {
            int currentNum = nums[i];

            int complement = target - currentNum;

            if (numMap.count(complement))
            {
                return {numMap[complement], i};
            }
            numMap[currentNum] = i;
        }
        return {};
    }
};

int main()
{
    Solution solution;

    std::vector<int> nums1 = {2, 7, 11, 15};
    int target1 = 9;
    std::vector<int> nums2 = {3, 2, 4};
    int target2 = 6;
    std::vector<int> nums3 = {3, 3};
    int target3 = 6;

    std::cout << "Teste 1:" << std::endl;
    std::vector<int> result1 = solution.twoSum(nums1, target1);
    std::cout << "Input: nums = [2,7,11,15], target = 9" << std::endl;
    std::cout << "Output: [" << result1[0] << ", " << result1[1] << "]" << std::endl;
    std::cout << "--------------------" << std::endl;

    std::cout << "Teste 2:" << std::endl;
    std::vector<int> result2 = solution.twoSum(nums2, target2);
    std::cout << "Input: nums = [3,2,4], target = 6" << std::endl;
    std::cout << "Output: [" << result2[0] << ", " << result2[1] << "]" << std::endl;
    std::cout << "--------------------" << std::endl;

    std::cout << "Teste 3:" << std::endl;
    std::vector<int> result3 = solution.twoSum(nums3, target3);
    std::cout << "Input: nums = [3,3], target = 6" << std::endl;
    std::cout << "Output: [" << result3[0] << ", " << result3[1] << "]" << std::endl;
    std::cout << "--------------------" << std::endl;


    return 0;
}