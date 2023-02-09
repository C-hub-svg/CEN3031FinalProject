import { OperationContext } from '@vercel/commerce/api/operations'
import { Category, GetSiteInfoOperation } from '@vercel/commerce/types/site'
import { LocalConfig } from '../index'

export type GetSiteInfoResult<
  T extends { categories: any[]; brands: any[] } = {
    categories: Category[]
    brands: any[]
  }
> = T

export default function getSiteInfoOperation({}: OperationContext<any>) {
  function getSiteInfo({
    query,
    variables,
    config: cfg,
  }: {
    query?: string
    variables?: any
    config?: Partial<LocalConfig>
    preview?: boolean
  } = {}): Promise<GetSiteInfoResult> {
    return Promise.resolve({
      categories: [
        {
          id: 'accesories',
          name: 'Accesories',
          slug: 'accesories',
          path: '/accesories',
        },
        {
          id: 'dresses',
          name: 'Dresses',
          slug: 'dresses',
          path: '/dresses',
        },
        {
          id: 'jackets',
          name: 'Jackets',
          slug: 'jackets',
          path: '/jackets',
        },
        {
          id: 'pants',
          name: 'Pants',
          slug: 'pants',
          path: '/pants',
        },
        {
          id: 'shorts',
          name: 'Shorts',
          slug: 'shorts',
          path: '/shorts',
        },
        {
          id: 'skirts',
          name: 'Skirts',
          slug: 'skirts',
          path: '/skirts',
        },
        
        {
          id: 'tops',
          name: 'Tops',
          slug: 'tops',
          path: '/tops',
        },
        {
          id: 'business',
          name: 'Business',
          slug: 'business',
          path: '/business',
        },
        {
          id: 'formal',
          name: 'Formal',
          slug: 'formal',
          path: '/formal',
        },
        
      ],
      brands: [
        {
          id: 'forever21',
          name: 'Forever21',
          slug: 'forever21',
          path: '/forever21',
        },
        {
          id: 'nike',
          name: 'Nike',
          slug: 'nike',
          path: '/nike',
        },
        {
          id: 'shein',
          name: 'Shein',
          slug: 'shein',
          path: '/shein',
        },
      ],
    })
  }

  return getSiteInfo
}
